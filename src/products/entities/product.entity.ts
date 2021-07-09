import { Category } from "src/categories/entities/category.entity";
import { Wish } from "src/wishes/entities/wish.entity";
import { BaseEntity, Column, Entity, Index, ManyToOne, OneToMany, PrimaryGeneratedColumn } from "typeorm";

@Entity('products')
export class Product extends BaseEntity {
  @Index({ unique: true })
  @PrimaryGeneratedColumn()
  id: number;

  @Index({ fulltext: true })
  @Column('text')
  title: string;

  @Column('text')
  description: string;

  @Index({ unique: true })
  @Column()
  productKey: string;

  @Column('text')
  imageUrl: string;

  @Index()
  @Column('double')
  rating: number;

  @Index()
  @Column('double')
  price: number;

  @Column()
  currency: string;

  @Column('datetime')
  modified: Date = new Date(Date.now());

  @Index()
  @ManyToOne(() => Category, category => category.products, { onDelete: 'CASCADE' })
  category: Category;

  @Column('text')
  website: string;

  @OneToMany(() => Wish, wish => wish.product)
  wishes: Wish[]
}
