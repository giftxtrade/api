import { Entity, BaseEntity, PrimaryColumn, PrimaryGeneratedColumn, Column } from 'typeorm';

@Entity('participants')
export class Participant extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;
}
