import { Entity, BaseEntity, PrimaryGeneratedColumn, Column, ManyToOne, ManyToMany } from 'typeorm';
import { User } from 'src/users/entities/user.entity';
import { Wish } from 'src/wishes/entities/wish.entity';
import { Event } from 'src/events/entities/event.entity';
import { Draw } from 'src/draws/entities/draw.entity';

@Entity('participants')
export class Participant extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Column()
  email: string;

  @Column()
  address: string;

  @Column('boolean')
  organizer: boolean = false;

  @Column('boolean')
  participates: boolean = true;

  @Column('boolean')
  accepted: boolean = false;

  @ManyToOne(() => Event, event => event.participants)
  event: Event;

  @ManyToOne(() => User, user => user.participated)
  user: User;

  @ManyToMany(() => Wish, wish => wish.participant)
  wishes: Wish[];

  @ManyToMany(() => Draw, draw => draw.drawer)
  drawers: Draw[];

  @ManyToMany(() => Draw, draw => draw.drawee)
  drawees: Draw[];
}
