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
  ) { }

  async create(createParticipantDto: CreateParticipantDto, event: Event, user?: User): Promise<Participant> {
    const participant = new Participant();
    participant.name = createParticipantDto.name;
    participant.email = createParticipantDto.email;
    participant.address = createParticipantDto.address;
    participant.organizer = createParticipantDto.organizer;
    participant.participates = createParticipantDto.participates;
    participant.accepted = createParticipantDto.accepted;
    if (user)
      participant.user = user;
    participant.event = event;

    return await participant.save();
  }

  async findAll(): Promise<Participant[]> {
    return await this.participantRepository.find();
  }

  async findOne(id: number): Promise<Participant> {
    return await this.participantRepository.findOne({ id });
  }

  async getPendingParticipantForEvent(event: Event, user: User): Promise<Participant> {
    return await this.participantRepository
      .createQueryBuilder('p')
      .where('p.eventId = :eventId AND p.email = :email AND p.accepted = false', {
        eventId: event.id,
        email: user.email
      })
      .getOne();
  }

  async acceptEvent(event: Event, user: User): Promise<Participant> {
    const participant = await this.getPendingParticipantForEvent(event, user)
    if (!participant) {
      throw new HttpException({
        message: 'Operation failed'
      }, HttpStatus.BAD_REQUEST);
    }

    participant.accepted = true;
    participant.user = user;
    return await participant.save();
  }

  async declineEvent(event: Event, user: User): Promise<boolean> {
    const participant = await this.getPendingParticipantForEvent(event, user)
    if (!participant) {
      throw new HttpException({
        message: 'Operation failed'
      }, HttpStatus.BAD_REQUEST);
    }

    await this.participantRepository.delete({ id: participant.id });
    return true;
  }

  update(id: number, updateParticipantDto: UpdateParticipantDto) {
    return `This action updates a #${id} participant`;
  }

  remove(id: number) {
    return `This action removes a #${id} participant`;
  }
}
