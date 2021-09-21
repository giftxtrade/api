import { HttpException, Injectable, HttpStatus } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { Repository } from 'typeorm';
import { Event } from './entities/event.entity'
import { Participant } from 'src/participants/entities/participant.entity';
import { ParticipantsService } from '../participants/participants.service';
import { User } from 'src/users/entities/user.entity';
import { CreateParticipantDto } from 'src/participants/dto/create-participant.dto';
import Link from 'src/links/entity/link.entity';
import { LinksService } from '../links/links.service';

@Injectable()
export class EventsService {
  constructor(
    @InjectRepository(Event)
    private readonly eventsRepository: Repository<Event>,
    private readonly participantsService: ParticipantsService,
    private readonly linksService: LinksService,
  ) { }

  async create(createEventDto: CreateEventDto, organizer: User): Promise<Event> {
    const event = new Event();
    event.name = createEventDto.name;
    event.description = createEventDto.description;
    event.budget = createEventDto.budget;
    event.invitationMessage = createEventDto.invitationMessage;
    event.drawAt = createEventDto.drawAt;
    event.closeAt = createEventDto.closeAt;

    const newEvent = await event.save();

    const participants = await this.addAllParticipants(createEventDto.participants, organizer, newEvent);
    return await this.findOneForUser(newEvent.id, organizer);
  }

  async findOne(id: number): Promise<Event> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p')
      .where('e.id = :eventId', {
        eventId: id
      })
      .getOne();
  }

  async findOneForUser(id: number, user: User): Promise<Event> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p')
      .where('e.id = :eventId AND p.userId = :userId', {
        userId: user.id,
        eventId: id
      })
      .getOne();
  }

  async findOneForOrganizerUser(id: number, user: User): Promise<Event> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p')
      .where('e.id = :eventId AND p.userId = :userId AND p.organizer = true', {
        userId: user.id,
        eventId: id
      })
      .getOne();
  }

  async findAllForUserNoParticipants(user: User): Promise<Event[]> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoin('e.participants', 'p')
      .where('p.userId = :userId', { userId: user.id })
      .orderBy('e.drawAt', 'DESC')
      .getMany();
  }

  async findAllForUser(user: User): Promise<Event[]> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p')
      .where('p.userId = :userId', { userId: user.id })
      .orderBy('e.drawAt', 'DESC')
      .getMany();
  }

  async findAllForUserWithParticipantUser(user: User) {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p1')
      .leftJoinAndSelect('e.participants', 'p2')
      .where('p1.userId = :userId', { userId: user.id })
      .orderBy('e.drawAt', 'DESC')
      .getMany();
  }

  async findAllInvitesForUser(user: User): Promise<Event[]> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p')
      .where('p.accepted = 0 AND p.email = :email', {
        email: user.email
      })
      .getMany();
  }

  async createLinkForEvent(event: Event, user: User, expirationDate: Date): Promise<Link> {
    const participant = await this.participantsService
      .findByEventAndOrganizer(event, user);
    if (!participant) {
      throw new HttpException({
        message: 'Could not perform operation'
      }, HttpStatus.BAD_REQUEST)
    }
    return this.linksService.create(event, expirationDate);
  }

  async findOneByLink(link: Link): Promise<Event> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoin('e.links', 'l')
      .where('l.code = :linkCode', { linkCode: `${link.code}` })
      .getOne();
  }

  async isUserPartOfEvent(event: Event, user: User): Promise<boolean> {
    const participant = await this.participantsService.findByEventAndUser(event, user);
    return participant ? true : false;
  }

  async isUserPartOfEventShallow(event: Event, user: User): Promise<boolean> {
    const participant = await this.participantsService.findByEventAndShallowUser(event, user.email);
    return participant ? true : false;
  }

  async update(event: Event, updateEventDto: UpdateEventDto) {
    const updated = await this.eventsRepository.update({ id: event.id }, updateEventDto);

    // If draw date is updated then also update link expiration
    if (updateEventDto.drawAt) {
      // No need to use await since we don't return the value
      this.linksService.updateExpriationDate(event, updateEventDto.drawAt);
    }
    return await this.eventsRepository.findOne(event.id);
  }

  async remove(id: number) {
    return await this.eventsRepository.delete({ id: id });
  }

  private checkForMainOrganizer(createParticipants: CreateParticipantDto[], organizer: User): boolean {
    let found = false;
    createParticipants.forEach(p => {
      if (p.email === organizer.email) {
        found = true;

        if (!p.organizer) {
          throw new HttpException({
            message: `${organizer.name} (${organizer.email}) must have organizer set as \`true\``
          }, HttpStatus.BAD_REQUEST);
        }
      }
    });
    return found
  }

  private async addAllParticipants(createParticipants: CreateParticipantDto[], organizer: User, event: Event): Promise<Participant[]> {
    if (!this.checkForMainOrganizer(createParticipants, organizer)) {
      throw new HttpException({
        message: `${organizer.name} (${organizer.email}) must be a participant. If you don't want to participate set participates to \`false\``
      }, HttpStatus.BAD_REQUEST);
    }

    // Add all participants
    const participants = Array<Participant>();
    for (const p of createParticipants) {
      if (p.email === '' || p.name === '')
        continue;

      if (p.email === organizer.email) {
        // The main organizer must have to have a valid account. 
        // Therefore, set accepted to true
        p.accepted = true;
        participants.push(await this.participantsService.create(p, event, organizer));
      } else {
        participants.push(await this.participantsService.create(p, event));
      }
    }
    return participants;
  }
}
