import { Category } from "src/categories/entities/category.entity";
import { BaseEntity, Column, Entity, ManyToOne, OneToMany, PrimaryGeneratedColumn } from "typeorm";

@Entity('products')
export class Product extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  title: string;

  @Column('text')
  description: string;

  @Column()
  productKey: string;

  @Column()
  imageUrl: string;

  @Column('double')
  rating: number;

  @Column('double')
  price: number;

  @Column()
  currency: string;

  @Column('datetime')
  modified: Date = new Date(Date.now());

  @ManyToOne(() => Category, category => category.products)
  category: Category;

  @Column()
  website: string;
}
