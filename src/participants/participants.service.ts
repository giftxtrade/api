import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Event } from 'src/events/entities/event.entity';
import { User } from 'src/users/entities/user.entity';
import { Repository } from 'typeorm';
import { CreateParticipantDto } from './dto/create-participant.dto';
import { UpdateParticipantDto } from './dto/update-participant.dto';
import { Participant } from './entities/participant.entity';

@Injectable()
export class ParticipantsService {
  constructor(
    @InjectRepository(Participant)
    private readonly participantRepository: Repository<Participant>,
  ) {}

  async create(
    createParticipantDto: CreateParticipantDto,
    event: Event,
    user?: User,
  ): Promise<Participant> {
    if (user) {
      const participant = await this.findByEventAndUser(event, user);
      if (participant) {
        return participant;
      }
    } else {
      const shallowParticipant = await this.findByEventAndShallowUser(
        event,
        createParticipantDto.email,
      );
      if (shallowParticipant) {
        return shallowParticipant;
      }
    }

    const participant = new Participant();
    participant.name = createParticipantDto.name;
    participant.email = createParticipantDto.email;
    participant.address = createParticipantDto.address;
    participant.organizer = createParticipantDto.organizer;
    participant.participates = createParticipantDto.participates;
    participant.accepted = createParticipantDto.accepted;
    if (user) participant.user = user;
    participant.event = event;

    return await participant.save();
  }

  async findAll(): Promise<Participant[]> {
    return await this.participantRepository.find();
  }

  async findOne(id: number): Promise<Participant> {
    return await this.participantRepository.findOne({ id });
  }

  async findOneWithUser(id: number): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .innerJoinAndSelect('p.user', 'u')
      .where('p.id = :participantId', { participantId: id })
      .getOne();
  }

  /**
   * Joins event and the participant user
   * @param event
   * @param participantId
   * @returns
   */
  async findOneByEvent(
    event: Event,
    participantId: number,
  ): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .innerJoin('p.event', 'e')
      .innerJoinAndSelect('p.user', 'u')
      .where('e.id = :eventId AND p.id = :participantId', {
        eventId: event.id,
        participantId,
      })
      .getOne();
  }

  async findAllByEvent(event: Event): Promise<Participant[]> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .innerJoin('p.event', 'e')
      .where('e.id = :eventId', { eventId: event.id })
      .getMany();
  }

  async findAllByEventWithUser(event: Event): Promise<Participant[]> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .innerJoin('p.event', 'e')
      .leftJoinAndSelect('p.user', 'u')
      .where('e.id = :eventId', { eventId: event.id })
      .getMany();
  }

  async findByEventAndUser(event: Event, user: User): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .where('p.eventId = :eventId AND p.userId = :userId', {
        eventId: event.id,
        userId: user.id,
      })
      .getOne();
  }

  async findByEventAndShallowUser(
    event: Event,
    email: string,
  ): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .where('p.eventId = :eventId AND p.email = :email', {
        eventId: event.id,
        email: `${email}`,
      })
      .getOne();
  }

  async findByEventAndOrganizer(
    event: Event,
    user: User,
  ): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .innerJoin('p.event', 'e')
      .where('e.id = :eventId AND p.userId = :userId AND p.organizer = true', {
        eventId: event.id,
        userId: user.id,
      })
      .getOne();
  }

  async getPendingParticipantForEvent(
    event: Event,
    user: User,
  ): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .where(
        'p.eventId = :eventId AND p.email = :email AND p.accepted = false',
        {
          eventId: event.id,
          email: user.email,
        },
      )
      .getOne();
  }

  async acceptEvent(event: Event, user: User): Promise<Participant> {
    const participant = await this.getPendingParticipantForEvent(event, user);
    if (!participant || participant.accepted) {
      throw new HttpException(
        {
          message: 'Operation failed',
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    participant.accepted = true;
    participant.user = user;
    return await participant.save();
  }

  async declineEvent(event: Event, user: User): Promise<boolean> {
    const participant = await this.getPendingParticipantForEvent(event, user);
    if (!participant || participant.accepted) {
      throw new HttpException(
        {
          message: 'Operation failed',
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    await this.participantRepository.delete({ id: participant.id });
    return true;
  }

  async update(id: number, updateParticipantDto: UpdateParticipantDto) {
    return await this.participantRepository.update(
      { id },
      updateParticipantDto,
    );
  }

  async remove(id: number) {
    return await this.participantRepository.delete({ id });
  }
}
