import { Module } from '@nestjs/common';
import { EventsService } from './events.service';
import { EventsController } from './events.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Event } from './entities/event.entity';
import { ParticipantsModule } from 'src/participants/participants.module';
import { UsersModule } from 'src/users/users.module';
import { LinksModule } from 'src/links/links.module';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Event
    ]),
    ParticipantsModule,
    UsersModule,
    LinksModule
  ],
  controllers: [EventsController],
  providers: [EventsService],
  exports: [EventsService]
})
export class EventsModule {}
