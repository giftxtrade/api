import { Product } from "src/products/entities/product.entity";
import { BaseEntity, Column, Entity, OneToMany, PrimaryGeneratedColumn, Index } from 'typeorm';

@Entity('categories')
export class Category extends BaseEntity {
  @Index({ unique: true })
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @Column('text', { nullable: true })
  description: string;

  @Column()
  categoryUrl: string;

  @OneToMany(() => Product, product => product.category)
  products: Product[];
}
