import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpException, HttpStatus } from '@nestjs/common';
import { WishesService } from './wishes.service';
import { CreateWishDto } from './dto/create-wish.dto';
import { UpdateWishDto } from './dto/update-wish.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';
import { EventsService } from 'src/events/events.service';
import { ParticipantsService } from 'src/participants/participants.service';

@Controller('wishes')
export class WishesController {
  constructor(
    private readonly wishesService: WishesService,
    private readonly usersService: UsersService,
    private readonly eventsService: EventsService,
    private readonly participantService: ParticipantsService,
  ) { }

  @UseGuards(JwtAuthGuard)
  @Post()
  async create(@Request() req, @Body() createWishDto: CreateWishDto) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    return await this.wishesService.create(user, createWishDto);
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  async findAll(@Request() req, @Param('id') eventId) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const event = await this.eventsService.findOneForUser(eventId, user);
    if (!event) {
      throw new HttpException({
        message: 'Event not found'
      }, HttpStatus.NOT_FOUND);
    }

    return await this.wishesService.findAllByUserEvent(user, event);
  }

  @UseGuards(JwtAuthGuard)
  @Get(':eventId/:participantId')
  async findAllForParticipant(@Request() req, @Param('eventId') eventId, @Param('participantId') participantId) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    const participant = await this.participantService.findOneWithUser(participantId);
    const event = await this.eventsService.findOneForUser(eventId, user);

    if (!participant || !event) {
      throw new HttpException({
        message: 'Invalid participant or event'
      }, HttpStatus.NOT_FOUND);
    }
    return await this.wishesService.findAllByUserEvent(participant.user, event);
  }

  @UseGuards(JwtAuthGuard)
  @Delete()
  async remove(@Request() req, @Body() createWishDto: CreateWishDto) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    return await this.wishesService.remove(user, createWishDto);
  }
}
