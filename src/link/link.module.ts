import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { LinkService } from './link.service';
import Link from './entity/link.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Link
    ])
  ],
  providers: [LinkService],
  exports: [LinkService]
})
export class LinkModule { }
