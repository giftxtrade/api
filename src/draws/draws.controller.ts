import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpException, HttpStatus } from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { EventsService } from 'src/events/events.service';
import { ParticipantsService } from 'src/participants/participants.service';
import { UsersService } from 'src/users/users.service';
import { DrawsService } from './draws.service';

@Controller('draws')
export class DrawsController {
  constructor(
    private readonly drawsService: DrawsService,
    private readonly usersService: UsersService,
    private readonly eventsService: EventsService,
    private readonly participantsService: ParticipantsService,
  ) { }

  @UseGuards(JwtAuthGuard)
  @Post()
  async create(@Request() res, @Body() body: { eventId: number }) {
    const user = await this.usersService.findByEmail(res.user.user.email);
    const event = await this.eventsService.findOneForOrganizerUser(body.eventId, user);
    if (!event) {
      throw new HttpException({
        message: 'Something went wrong'
      }, HttpStatus.BAD_REQUEST);
    }
    return await this.drawsService.create(event, user);
  }

  @UseGuards(JwtAuthGuard)
  @Get(':eventId')
  async findAll(@Request() res, @Param('eventId') eventId: number) {
    const user = await this.usersService.findByEmail(res.user.user.email);
    const event = await this.eventsService.findOneForOrganizerUser(eventId, user);
    if (!event) {
      throw new HttpException({
        message: 'Something went wrong'
      }, HttpStatus.BAD_REQUEST);
    }
    return this.drawsService.findAll(event);
  }

  @UseGuards(JwtAuthGuard)
  @Get('me/:eventId')
  async findForMe(@Request() res, @Param('eventId') eventId: number) {
    const user = await this.usersService.findByEmail(res.user.user.email);
    const event = await this.eventsService.findOneForUser(eventId, user);
    if (!event) {
      throw new HttpException({
        message: 'Something went wrong'
      }, HttpStatus.BAD_REQUEST);
    }

    const participant = await this.participantsService.findByEventAndUser(event, user);
    if (!participant) {
      throw new HttpException({
        message: 'Something went wrong'
      }, HttpStatus.BAD_REQUEST);
    }

    return this.drawsService.findForParticipant(event, participant);
  }
}
