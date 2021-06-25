import { Module } from '@nestjs/common';
import { DrawsService } from './draws.service';
import { DrawsController } from './draws.controller';

@Module({
  controllers: [DrawsController],
  providers: [DrawsService]
})
export class DrawsModule {}
