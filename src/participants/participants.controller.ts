import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpStatus, HttpException, Query } from '@nestjs/common';
import { ParticipantsService } from './participants.service';
import { CreateParticipantDto } from './dto/create-participant.dto';
import { UpdateParticipantDto } from './dto/update-participant.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';
import { EventsService } from 'src/events/events.service';

@Controller('participants')
export class ParticipantsController {
  constructor(
    private readonly participantsService: ParticipantsService,
    private readonly usersService: UsersService,
    private readonly eventsService: EventsService,
  ) { }

  @UseGuards(JwtAuthGuard)
  @Patch(':participantId')
  async update(@Request() req, @Param('participantId') participantId: number, @Body() { address }: { address: string }) {
    const user = await this.usersService.findByEmail(req.user.user.email);

    const participant = await this.participantsService.findOneWithUser(participantId);
    if (!participant) {
      throw new HttpException({
        message: 'Participant does not exist'
      }, HttpStatus.NOT_FOUND);
    }
    if (participant.user.id !== user.id) {
      throw new HttpException({
        message: 'Could not update address'
      }, HttpStatus.BAD_REQUEST);
    }

    participant.address = address;
    return await participant.save();
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':participantId')
  async remove(@Request() req, @Param('participantId') participantId: number) {
    const user = await this.usersService.findByEmail(req.user.user.email);

    const participant = await this.participantsService.findOneWithUser(participantId);
    if (!participant) {
      throw new HttpException({
        message: 'Participant does not exist'
      }, HttpStatus.NOT_FOUND);
    }
    if (participant.user.id !== user.id || participant.organizer) {
      throw new HttpException({
        message: 'Could not delete'
      }, HttpStatus.BAD_REQUEST);
    }
    return await this.participantsService.remove(participantId);
  }

  @UseGuards(JwtAuthGuard)
  @Delete('manage')
  async organizerRemove(@Request() req, @Query('participantId') participantId: number, @Query('eventId') eventId: number) {
    const organizerUser = await this.usersService.findByEmail(req.user.user.email);

    // Find event
    const event = await this.eventsService.findOne(eventId);
    if (!event) {
      throw new HttpException({
        message: "Event not found"
      }, HttpStatus.NOT_FOUND);
    }

    // Get auth user as participant and check if they are an organizer 
    const organizer = await this.participantsService.findByEventAndUser(event, organizerUser);
    if (!organizer || !organizer?.organizer) {
      throw new HttpException({
        message: "Illegal action"
      }, HttpStatus.BAD_REQUEST);
    }

    const participant = await this.participantsService.findOneWithUser(participantId);
    if (!participant) {
      throw new HttpException({
        message: 'Participant does not exist'
      }, HttpStatus.BAD_REQUEST);
    }
    if (participant?.event.id !== event.id) {
      throw new HttpException({
        message: 'Could not remove participant'
      }, HttpStatus.BAD_REQUEST);
    }

    const removedParticipant = await this.participantsService.remove(participantId);
    return { message: 'Participant removed' }
  }
}
