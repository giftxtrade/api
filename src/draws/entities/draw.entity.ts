import { Event } from 'src/events/entities/event.entity';
import { Participant } from 'src/participants/entities/participant.entity';
import { BaseEntity, Column, Entity, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';

@Entity('draws')
export class Draw extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column('datetime')
  createAt: Date = new Date(Date.now());

  @ManyToOne(() => Participant, participant => participant.drawers, { onDelete: 'CASCADE' })
  drawer: Participant;

  @ManyToOne(() => Participant, participant => participant.drawees, { onDelete: 'CASCADE' })
  drawee: Participant;

  @ManyToOne(() => Event, event => event.draws, { onDelete: 'CASCADE' })
  event: Event;
}
