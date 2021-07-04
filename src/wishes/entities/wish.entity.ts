import { Entity, BaseEntity, PrimaryGeneratedColumn, Column, ManyToOne } from 'typeorm';
import { User } from 'src/users/entities/user.entity';
import { Participant } from 'src/participants/entities/participant.entity';
import { Product } from 'src/products/entities/product.entity';
import { Event } from 'src/events/entities/event.entity';

@Entity('wishes')
export class Wish extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column('datetime')
  createdAt: Date = new Date(Date.now());

  @ManyToOne(() => User, user => user.wishes, { onDelete: 'CASCADE' })
  user: User;

  @ManyToOne(() => Participant, participant => participant.wishes, { onDelete: 'CASCADE' })
  participant: Participant;

  @ManyToOne(() => Product, product => product.wishes, { onDelete: 'CASCADE' })
  product: Product;

  @ManyToOne(() => Event, event => event.wishes, { onDelete: 'CASCADE' })
  event: Event;
}
