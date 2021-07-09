import { BaseEntity, Column, Entity, PrimaryGeneratedColumn, ManyToOne, Index } from 'typeorm';
import { Event } from 'src/events/entities/event.entity';

@Entity('links')
export default class Link extends BaseEntity {
  @Index({ unique: true })
  @PrimaryGeneratedColumn()
  id: number;

  @Index({ unique: true })
  @Column({ unique: true })
  code: string;

  @Column('datetime')
  createdAt: Date = new Date(Date.now());

  @Column('datetime')
  expirationDate: Date;

  @Index()
  @ManyToOne(() => Event, event => event.links, { onDelete: 'CASCADE' })
  event: Event;
}