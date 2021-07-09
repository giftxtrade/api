import { Wish } from "src/wishes/entities/wish.entity";
import { BaseEntity, Column, Entity, Index, PrimaryGeneratedColumn } from "typeorm";
import { Participant } from 'src/participants/entities/participant.entity';

@Entity('users')
export class User extends BaseEntity {
  @Index({ unique: true })
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Index({ unique: true })
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
