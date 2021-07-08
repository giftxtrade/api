import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, Query } from '@nestjs/common';
import { ParticipantsService } from './participants.service';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';
import { EventsService } from 'src/events/events.service';
import { BAD_REQUEST, NOT_FOUND } from 'src/util/exceptions';
import { UpdateParticipantDto } from './dto/update-participant.dto';

@Controller('participants')
export class ParticipantsController {
  constructor(
    private readonly participantsService: ParticipantsService,
    private readonly usersService: UsersService,
    private readonly eventsService: EventsService,
  ) { }

  @UseGuards(JwtAuthGuard)
  @Delete('manage')
  async organizerRemove(@Request() req, @Query('participantId') participantId: number, @Query('eventId') eventId: number) {
    const { event, participant, organizerUser } = await this.validateEventAndParticipant(req.user.user.email, eventId, participantId);

    const shallowParticipant = await this.participantsService.findByEventAndShallowUser(event, participant.email)
    if (!shallowParticipant || participant.email === organizerUser.email)
      throw BAD_REQUEST('Could not remove participant');

    await this.participantsService.remove(participantId);
    return { message: 'Participant removed' };
  }

  @UseGuards(JwtAuthGuard)
  @Patch('manage')
  async changeParticipantOrganizerStatus(
    @Request() req,
    @Query('participantId') participantId: number,
    @Query('eventId') eventId: number,
    @Body() { organizer }: { organizer: boolean }
  ) {
    const { event, participant, organizerUser } = await this.validateEventAndParticipant(req.user.user.email, eventId, participantId);

    const shallowParticipant = await this.participantsService.findByEventAndShallowUser(event, participant.email)
    if (!shallowParticipant || participant.email === organizerUser.email)
      throw BAD_REQUEST('Could not update participant');

    await this.participantsService.update(participantId, { organizer: organizer });
    return await this.participantsService.findOne(participantId);
  }

  @UseGuards(JwtAuthGuard)
  @Patch(':participantId')
  async updateParticipantAddress(@Request() req, @Param('participantId') participantId: number, @Body() { address }: { address: string }) {
    const user = await this.usersService.findByEmail(req.user.user.email);

    const participant = await this.participantsService.findOneWithUser(participantId);
    if (!participant)
      throw NOT_FOUND('Participant does not exist');
    if (participant.user.id !== user.id)
      throw BAD_REQUEST('Could not update address');

    if (participant.address === address)
      return participant;

    participant.address = address;
    return await participant.save();
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':participantId')
  async remove(@Request() req, @Param('participantId') participantId: number) {
    const user = await this.usersService.findByEmail(req.user.user.email);

    const participant = await this.participantsService.findOneWithUser(participantId);
    if (!participant)
      throw NOT_FOUND('Participant does not exist');
    if (participant.user.id !== user.id || participant.organizer)
      throw BAD_REQUEST('Could not delete');
    await this.participantsService.remove(participantId)
    return { message: 'Removed from event' };
  }

  private async validateEventAndParticipant(email: string, eventId: number, participantId: number) {
    const organizerUser = await this.usersService.findByEmail(email);

    // Find event
    const event = await this.eventsService.findOne(eventId);
    if (!event) {
      throw NOT_FOUND("Event not found");
    }

    // Get auth user as participant and check if they are an organizer 
    const organizer = await this.participantsService.findByEventAndUser(event, organizerUser);
    if (!organizer || !organizer?.organizer)
      throw BAD_REQUEST("Illegal action");

    const participant = await this.participantsService.findOne(participantId);
    if (!participant)
      throw NOT_FOUND('Participant does not exist');

    return { event, organizerUser, organizer, participant };
  }
}
