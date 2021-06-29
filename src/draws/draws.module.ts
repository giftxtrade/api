import { Module } from '@nestjs/common';
import { DrawsService } from './draws.service';
import { DrawsController } from './draws.controller';
import { ParticipantsModule } from 'src/participants/participants.module';
import { EventsModule } from 'src/events/events.module';
import { UsersModule } from 'src/users/users.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Draw } from './entities/draw.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Draw
    ]),
    ParticipantsModule,
    EventsModule,
    UsersModule,
  ],
  controllers: [DrawsController],
  providers: [DrawsService],
  exports: [DrawsService]
})
export class DrawsModule {}
