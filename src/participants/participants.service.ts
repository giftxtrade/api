import { Injectable } from '@nestjs/common';
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

  update(id: number, updateParticipantDto: UpdateParticipantDto) {
    return `This action updates a #${id} participant`;
  }

  remove(id: number) {
    return `This action removes a #${id} participant`;
  }
}
