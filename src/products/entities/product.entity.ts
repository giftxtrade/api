import { Category } from "src/categories/entities/category.entity";
import { Wish } from "src/wishes/entities/wish.entity";
import { BaseEntity, Column, Entity, ManyToOne, OneToMany, PrimaryGeneratedColumn } from "typeorm";

@Entity('products')
export class Product extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column('text')
  title: string;

  @Column('text')
  description: string;

  @Column()
  productKey: string;

  @Column('text')
  imageUrl: string;

  @Column('double')
  rating: number;

  @Column('double')
  price: number;

  @Column()
  currency: string;

  @Column('datetime')
  modified: Date = new Date(Date.now());

  @ManyToOne(() => Category, category => category.products, { onDelete: 'CASCADE' })
  category: Category;

  @Column('text')
  website: string;

  @OneToMany(() => Wish, wish => wish.product)
  wishes: Wish[]
}
