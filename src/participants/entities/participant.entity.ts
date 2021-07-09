import { Entity, BaseEntity, PrimaryGeneratedColumn, Column, ManyToOne, ManyToMany, Index } from 'typeorm';
import { User } from 'src/users/entities/user.entity';
import { Wish } from 'src/wishes/entities/wish.entity';
import { Event } from 'src/events/entities/event.entity';
import { Draw } from 'src/draws/entities/draw.entity';

@Entity('participants')
export class Participant extends BaseEntity {
  @Index({ unique: true })
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Index()
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

  @Index()
  @ManyToOne(() => Event, event => event.participants, { onDelete: 'CASCADE' })
  event: Event;

  @Index()
  @ManyToOne(() => User, user => user.participated, { onDelete: 'CASCADE' })
  user: User;

  @ManyToMany(() => Wish, wish => wish.participant)
  wishes: Wish[];

  @ManyToMany(() => Draw, draw => draw.drawer)
  drawers: Draw[];

  @ManyToMany(() => Draw, draw => draw.drawee)
  drawees: Draw[];
}
