import { Entity, BaseEntity, PrimaryColumn, PrimaryGeneratedColumn, Column, OneToOne } from 'typeorm';
import { User } from 'src/users/entities/user.entity';
import { Event } from 'src/events/entities/event.entity';
import { Participant } from 'src/participants/entities/participant.entity';
import { Product } from 'src/products/entities/product.entity';

@Entity('wishes')
export class Wish extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @OneToOne(() => User, user => user.wishes)
  user: User;

  @OneToOne(() => Participant, participant => participant.wishes)
  participant: Participant;

  @OneToOne(() => Product, product => product.wishes)
  product: Product;
}
