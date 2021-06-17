import { Module } from '@nestjs/common';
import { WishesService } from './wishes.service';
import { WishesController } from './wishes.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Wish } from './entities/wish.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Wish
    ])
  ],
  controllers: [WishesController],
  providers: [WishesService]
})
export class WishesModule {}
