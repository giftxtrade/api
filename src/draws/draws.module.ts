import { Module } from '@nestjs/common';
import { DrawsService } from './draws.service';
import { DrawsController } from './draws.controller';
import { ParticipantsModule } from 'src/participants/participants.module';
import { EventsService } from 'src/events/events.service';
import { UsersService } from 'src/users/users.service';

@Module({
  imports: [
    ParticipantsModule,
    EventsService,
    UsersService
  ],
  controllers: [DrawsController],
  providers: [DrawsService],
  exports: [DrawsService]
})
export class DrawsModule {}
