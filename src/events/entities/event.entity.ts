import { BaseEntity, Column, Entity, OneToMany, PrimaryGeneratedColumn } from 'typeorm';
import { Participant } from 'src/participants/entities/participant.entity';
import Link from 'src/links/entity/link.entity';
import { Wish } from 'src/wishes/entities/wish.entity';
import { Draw } from 'src/draws/entities/draw.entity';

@Entity('events')
export class Event extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Column('text', { nullable: true })
  description: string;

  @Column('decimal')
  budget: number;

  @Column('text')
  invitationMessage: string;

  @Column('datetime')
  createdAt: Date = new Date(Date.now());

  @Column('datetime')
  drawAt: Date;

  @Column('datetime')
  closeAt: Date;

  @OneToMany(() => Participant, participant => participant.event)
  participants: Participant[];

  @OneToMany(() => Link, link => link.event)
  links: Link[];

  @OneToMany(() => Wish, wish => wish.event)
  wishes: Wish[];

  @OneToMany(() => Draw, draw => draw.event)
  draws: Draw[];
}
