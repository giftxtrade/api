import { BaseEntity, Column, Entity, PrimaryGeneratedColumn, ManyToOne } from 'typeorm';
import { Event } from 'src/events/entities/event.entity';

@Entity('link')
export default class Link extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ unique: true })
  code: string;

  @Column('datetime')
  createdAt: Date = new Date(Date.now());

  @Column('datetime')
  expirationDate: Date;

  @ManyToOne(() => Event, event => event.links)
  event: Event;
}