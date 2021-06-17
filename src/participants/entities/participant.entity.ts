import { Entity, BaseEntity, PrimaryColumn, PrimaryGeneratedColumn, Column, OneToOne, OneToMany } from 'typeorm';
import { User } from 'src/users/entities/user.entity';
import { Wish } from 'src/wishes/entities/wish.entity';
import { Event } from 'src/events/entities/event.entity';

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

  @OneToOne(() => Event, event => event.participants)
  event: Event;

  @OneToOne(() => User, user => user.participated)
  user: User;

  @OneToMany(() => Wish, wish => wish.participant)
  wishes: Wish[];
}
