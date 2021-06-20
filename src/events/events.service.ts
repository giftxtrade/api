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

@Injectable()
export class EventsService {
  constructor(
    @InjectRepository(Event)
    private readonly eventsRepository: Repository<Event>,
    private readonly participantsService: ParticipantsService,
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

  findAll() {
    return `This action returns all events`;
  }

  async findOne(id: number): Promise<Event> {
    return await this.eventsRepository.findOne(id);
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

  async findAllForUser(user: User): Promise<Event[]> {
    return await this.eventsRepository
      .createQueryBuilder('e')
      .innerJoinAndSelect('e.participants', 'p')
      .where('p.userId = :userId', { userId: user.id })
      .orderBy('e.createdAt', 'DESC')
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

  update(id: number, updateEventDto: UpdateEventDto) {
    return `This action updates a #${id} event`;
  }

  remove(id: number) {
    return `This action removes a #${id} event`;
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
