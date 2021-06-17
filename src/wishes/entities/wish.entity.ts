import { Entity, BaseEntity, PrimaryGeneratedColumn, Column, ManyToOne } from 'typeorm';
import { User } from 'src/users/entities/user.entity';
import { Participant } from 'src/participants/entities/participant.entity';
import { Product } from 'src/products/entities/product.entity';

@Entity('wishes')
export class Wish extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne(() => User, user => user.wishes)
  user: User;

  @ManyToOne(() => Participant, participant => participant.wishes)
  participant: Participant;

  @ManyToOne(() => Product, product => product.wishes)
  product: Product;
}
