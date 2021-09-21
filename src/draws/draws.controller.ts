import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpException, HttpStatus } from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { EventsService } from 'src/events/events.service';
import { ParticipantsService } from 'src/participants/participants.service';
import { UsersService } from 'src/users/users.service';
import { DrawsService } from './draws.service';
import { BAD_REQUEST } from 'src/util/exceptions';
import { JwtStrategy } from '../auth/jwt.strategy';
import { sendMail, namesDrawnMail } from '../util/sendgrid';

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
      throw BAD_REQUEST('Something went wrong');
    }
    return await this.drawsService.create(event, user);
  }

  @UseGuards(JwtAuthGuard)
  @Get('confirm/:eventId')
  async confirmDraw(@Request() res, @Param('eventId') eventId: number) {
    const user = await this.usersService.findByEmail(res.user.user.email);
    const event = await this.eventsService.findOneForOrganizerUser(eventId, user);
    if (!event) {
      throw BAD_REQUEST('Something went wrong');
    }

    const draws = await this.drawsService.findAllWithUser(event);

    if (draws.length === 0)
      throw BAD_REQUEST('No draws found for event');

    // Mail all particiapnts about their draws
    draws.forEach(({ drawer, drawee }) => {
      namesDrawnMail(drawer.user, event, drawee);
    });
    return {
      message: "Participants are being notified!"
    };
  }

  @UseGuards(JwtAuthGuard)
  @Get(':eventId')
  async findAll(@Request() res, @Param('eventId') eventId: number) {
    const user = await this.usersService.findByEmail(res.user.user.email);
    const event = await this.eventsService.findOneForOrganizerUser(eventId, user);
    if (!event) {
      throw BAD_REQUEST('Something went wrong');
    }
    return await this.drawsService.findAllWithUser(event);
  }

  @UseGuards(JwtAuthGuard)
  @Get('me/:eventId')
  async findForMe(@Request() res, @Param('eventId') eventId: number) {
    const user = await this.usersService.findByEmail(res.user.user.email);
    const event = await this.eventsService.findOneForUser(eventId, user);
    if (!event) {
      throw BAD_REQUEST('Event not found');
    }

    const participant = await this.participantsService.findByEventAndUser(event, user);
    if (!participant) {
      throw BAD_REQUEST('Something went wrong');
    }

    const draw = await this.drawsService.findForParticipant(event, participant);
    if (!draw) {
      throw BAD_REQUEST('No draws found for user');
    }
    return draw;
  }
}
