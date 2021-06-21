import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpStatus, HttpException } from '@nestjs/common';
import { EventsService } from './events.service';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';
import { Event } from './entities/event.entity';
import { ParticipantsService } from '../participants/participants.service';
import Link from 'src/links/entity/link.entity';
import { LinksService } from 'src/links/links.service';

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
  async findAll(@Request() req): Promise<Event[]> {
    const user = await this.usersService.findByEmail(req.user.user.email);
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
      throw new HttpException({
        message: 'Event not found'
      }, HttpStatus.NOT_FOUND);
    }

    const participant = await this.participantsService.acceptEvent(event, user)
    return await this.eventsService.findOneForUser(eventId, user);
  }

  @UseGuards(JwtAuthGuard)
  @Get('/invites/decline/:eventId')
  async declineInvite(@Request() req, @Param('eventId') eventId: number): Promise<boolean> {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOne(eventId);
    if (!event) {
      throw new HttpException({
        message: 'Event not found'
      }, HttpStatus.NOT_FOUND);
    }

    return await this.participantsService.declineEvent(event, user);
  }

  @UseGuards(JwtAuthGuard)
  @Get(':eventId')
  async findOne(@Request() req, @Param('eventId') eventId: number): Promise<Event> {
    const user = await this.usersService.findByEmail(req.user.user.email);
    return await this.eventsService.findOneForUser(eventId, user);
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
      throw new HttpException({
        message: 'Invalid or expired invitation code.'
      }, HttpStatus.NOT_FOUND);
    }
    return link;
  }

  @UseGuards(JwtAuthGuard)
  @Get('invite-code/:inviteCode')
  async inviteCode(@Request() req, @Param('inviteCode') inviteCode: string) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const link = await this.linksService.findOne(inviteCode);
    if (!link) {
      throw new HttpException({
        message: 'Invalid or expired invitation code.'
      }, HttpStatus.NOT_FOUND);
    }
    const event = await this.eventsService.findOneByLink(link);
    const participant = await this.participantsService.create({
      name: user.name,
      email: user.email,
      address: '',
      participates: true,
      organizer: false,
      accepted: true
    }, event, user);
    return await this.eventsService.findOneForUser(event.id, user);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateEventDto: UpdateEventDto) {
    return this.eventsService.update(+id, updateEventDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.eventsService.remove(+id);
  }
}
