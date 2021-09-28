import { forwardRef, Module } from '@nestjs/common';
import { ParticipantsService } from './participants.service';
import { ParticipantsController } from './participants.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Participant } from './entities/participant.entity';
import { UsersModule } from 'src/users/users.module';
import { EventsModule } from 'src/events/events.module';
import { WishesModule } from '../wishes/wishes.module';
import { WishesService } from 'src/wishes/wishes.service';

@Module({
  imports: [
    TypeOrmModule.forFeature([Participant]),
    UsersModule,
    WishesModule,
    forwardRef(() => EventsModule),
  ],
  controllers: [ParticipantsController],
  providers: [ParticipantsService],
  exports: [ParticipantsService],
})
export class ParticipantsModule {}
