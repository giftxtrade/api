import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Draw } from './entities/draw.entity';
import { Repository } from 'typeorm';
import { Participant } from 'src/participants/entities/participant.entity';
import { ParticipantsService } from 'src/participants/participants.service';
import { User } from 'src/users/entities/user.entity';
import { Event } from 'src/events/entities/event.entity';
import { shuffle } from 'src/util/shuffle';

@Injectable()
export class DrawsService {
  constructor(
    @InjectRepository(Draw)
    private readonly drawsRepository: Repository<Draw>,
    private readonly participantsService: ParticipantsService
  ) { }

  async create(event: Event, user: User) {
    const allParticipants = await this.participantsService.findAllByEvent(event);
    const participants = shuffle(
      allParticipants
        .filter(p => p.accepted && p.participates)
    );
    return await this.generateDraw(participants, event);
  }

  async findAll(event: Event): Promise<Draw[]> {
    return await this.drawsRepository
      .createQueryBuilder('d')
      .leftJoinAndSelect('d.drawer', 'p1')
      .leftJoinAndSelect('d.drawee', 'p2')
      .where('d.eventId = :eventId', { eventId: event.id })
      .getMany();
  }

  async findForParticipant(event: Event, participant: Participant): Promise<Draw[]> {
    return await this.drawsRepository
      .createQueryBuilder('d')
      .leftJoinAndSelect('d.drawer', 'p1')
      .leftJoinAndSelect('d.drawee', 'p2')
      .where('d.eventId = :eventId AND d.drawerId = :drawerId',
        {
          eventId: event.id,
          drawerId: participant.id
        })
      .getMany();
  }

  private async generateDraw(participants: Participant[], event: Event) {
    await this.drawsRepository.createQueryBuilder()
      .where('eventId = :eventId', { eventId: event.id })
      .delete()
      .execute();

    const allDraws = this.getDraws(participants.length);
    const randomDraw = allDraws[Math.floor(Math.random() * allDraws.length)];
    const draws: Array<Draw> = [];

    for (let [key, val] of randomDraw) {
      const drawer = participants[key - 1];
      const drawee = participants[val - 1];

      const draw = new Draw();
      draw.event = event;
      draw.drawer = drawer;
      draw.drawee = drawee;
      draws.push(await draw.save());
    }
    return draws;
  }

  private getDraws(N: number): Array<Map<number, number>> {
    const allParticipantsMaps = [];

    for (let i = 0; i < N; i++) {
      const drawn = new Set<number>();
      const pairs = new Map<number, number>();
      const start = i + 1;

      for (let j = 0; j < N; j++) {
        const index = j + 1;
        let draw = start + index;
        if (draw == index)
          draw++;
        if (draw > N)
          draw %= N;
        if (draw === 0)
          draw++;

        if (drawn.has(draw) || draw === index) {
          pairs.clear();
          break;
        }
        drawn.add(draw);
        pairs.set(index, draw);
      }
      if (pairs.size == 0)
        continue;
      allParticipantsMaps.push(pairs);
    }
    return allParticipantsMaps;
  }
}
