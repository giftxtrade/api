import { BaseEntity, Column, PrimaryGeneratedColumn } from "typeorm";

export class User extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Column()
  email: string;

  @Column()
  image: string;

  @Column()
  phone: string;

  @Column('text', { select: false })
  password: string;

  @Column('boolean', { default: false })
  active: boolean;
}
