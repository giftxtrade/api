import { Wish } from "src/wishes/entities/wish.entity";
import { BaseEntity, Column, Entity, PrimaryGeneratedColumn } from "typeorm";
import { Participant } from 'src/participants/entities/participant.entity';

@Entity('users')
export class User extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Column()
  email: string;

  @Column()
  imageUrl: string;

  @Column({ nullable: true })
  phone: string;

  @Column('text', { select: false, nullable: true })
  password: string;

  participated: Participant[];

  wishes: Wish[];
}
