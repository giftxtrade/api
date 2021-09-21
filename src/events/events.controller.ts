import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpStatus, HttpException, Query } from '@nestjs/common';
import { EventsService } from './events.service';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';
import { Event } from './entities/event.entity';
import { ParticipantsService } from '../participants/participants.service';
import Link from 'src/links/entity/link.entity';
import { LinksService } from 'src/links/links.service';
import { Participant } from 'src/participants/entities/participant.entity';
import { BAD_REQUEST, NOT_FOUND } from 'src/util/exceptions';
import { newParticipantMail } from '../util/sendgrid';
import { User } from 'src/users/entities/user.entity';
import { FRONTEND_BASE } from '../../auth-tokens.json'

@Controller('events')
export class EventsController {
  constructor(
    private readonly eventsService: EventsService,
    private readonly usersService: UsersService,
    private readonly participantsService: ParticipantsService,
    private readonly linksService: LinksService,
  ) { }

  @UseGuards(JwtAuthGuard)
  @Post()
  async create(@Request() req, @Body() createEventDto: CreateEventDto): Promise<Event> {
    const user = await this.usersService.findOne(req.user.user.email);
    return await this.eventsService
      .create(createEventDto, user);
  }

  @UseGuards(JwtAuthGuard)
  @Get()
  async findAll(@Request() req, @Query('user') u: boolean): Promise<Event[]> {
    const user = await this.usersService.findByEmail(req.user.user.email);

    if (u)
      return await this.eventsService.findAllForUserWithParticipantUser(user);
    return await this.eventsService.findAllForUser(user);
  }

  @UseGuards(JwtAuthGuard)
  @Get('/invites')
  async findAllInvites(@Request() req): Promise<Event[]> {
    const user = await this.usersService.findByEmail(req.user.user.email);
    return await this.eventsService.findAllInvitesForUser(user);
  }

  @UseGuards(JwtAuthGuard)
  @Get('/invites/accept/:eventId')
  async acceptInvite(@Request() req, @Param('eventId') eventId: number): Promise<Event> {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOne(eventId);
    if (!event) {
      throw NOT_FOUND('Event not found');
    }

    const newParticipant = await this.participantsService.acceptEvent(event, user);
    const finalEvent = await this.eventsService.findOneForUser(eventId, user);

    // Send mail to all participants - newParticipant
    const allParticipants = await this.participantsService.findAllByEventWithUser(event);
    newParticipant.user = user; // Set user field so template can access imageUrl
    allParticipants.filter(p => p.email != newParticipant.email).forEach(p => {
      if (!p.accepted) {
        // Create fake user data if participant has not accepted invite yet
        const fakeUser = new User();
        fakeUser.imageUrl = `${FRONTEND_BASE}default.jpg`
        fakeUser.email = p.email;
        fakeUser.name = p.name;
        newParticipantMail(fakeUser, event, newParticipant);
      } else {
        newParticipantMail(p.user, event, newParticipant);
      }
    })

    return finalEvent;
  }

  @UseGuards(JwtAuthGuard)
  @Get('/invites/decline/:eventId')
  async declineInvite(@Request() req, @Param('eventId') eventId: number): Promise<boolean> {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOne(eventId);
    if (!event) {
      throw NOT_FOUND('Event not found');
    }

    return await this.participantsService.declineEvent(event, user);
  }

  @UseGuards(JwtAuthGuard)
  @Post('get-link/:eventId')
  async createLink(@Request() req, @Param('eventId') eventId: number, @Body() { expirationDate }: { expirationDate: Date }): Promise<Link> {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOne(eventId);
    return await this.eventsService
      .createLinkForEvent(event, user, expirationDate);
  }

  @Get('verify-invite-code/:inviteCode')
  async verifyInviteCode(@Param('inviteCode') inviteCode: string) {
    const link = await this.linksService.findOne(inviteCode);
    if (!link) {
      throw NOT_FOUND('Invalid or expired invitation code.');
    }
    return link;
  }

  @UseGuards(JwtAuthGuard)
  @Get('invite-code/:inviteCode')
  async inviteCode(@Request() req, @Param('inviteCode') inviteCode: string) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const link = await this.linksService.findOne(inviteCode);
    if (!link) {
      throw NOT_FOUND('Invalid or expired invitation code.');
    }
    const event = await this.eventsService.findOneByLink(link);
    const participant = await this.participantsService.create({
      name: user.name,
      email: user.email,
      address: '',
      participates: true,
      organizer: false,
      accepted: false
    }, event);
    return await this.eventsService.findOneForUser(event.id, user);
  }

  @UseGuards(JwtAuthGuard)
  @Get(':eventId')
  async findOne(@Request() req, @Param('eventId') eventId: number, @Query('verify') verify: boolean) {
    const user = await this.usersService.findByEmail(req.user.user.email);

    if (verify) {
      try {
        const eventInfo = await this.eventsService.findEventDetails(eventId, user);
        return {
          id: eventInfo.id,
          name: eventInfo.name,
          description: eventInfo.description
        }
      } catch (e) {
        throw NOT_FOUND("Event not found");
      }
    }

    try {
      return await this.eventsService.findOne(eventId, user);
    } catch (e) {
      throw NOT_FOUND("Event not found");
    }
  }

  @UseGuards(JwtAuthGuard)
  @Patch(':eventId')
  async update(@Request() req, @Param('eventId') eventId: number, @Body() updateEventDto: UpdateEventDto) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOneForUser(eventId, user);
    if (!event) {
      throw NOT_FOUND("Event not found");
    }

    const participant = await this.participantsService.findByEventAndOrganizer(event, user);
    if (!participant) {
      throw BAD_REQUEST("Operation not allowed for non-organizer users");
    }
    return await this.eventsService.update(event, updateEventDto);
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':eventId')
  async remove(@Request() req, @Param('eventId') eventId: number) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOneForUser(eventId, user);
    if (!event) {
      throw NOT_FOUND("Event not found");
    }

    const participant = await this.participantsService.findByEventAndOrganizer(event, user);
    if (!participant) {
      throw BAD_REQUEST("Delete not allowed for non-organizer users");
    }
    const deleteStatus = await this.eventsService.remove(eventId);
    return {
      message: 'Event deleted'
    };
  }

  @Get('get-details/:linkCode')
  async getEventDetailsFromCode(@Param('linkCode') linkCode: string): Promise<{ name: string, description: string }> {
    const link = await this.linksService.findOneWithEvent(linkCode.trim());
    if (!link)
      throw BAD_REQUEST('Invalid invite code');

    return {
      name: link.event.name,
      description: link.event.description
    };
  }
}
