import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpException, HttpStatus } from '@nestjs/common';
import { WishesService } from './wishes.service';
import { CreateWishDto } from './dto/create-wish.dto';
import { UpdateWishDto } from './dto/update-wish.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';
import { EventsService } from 'src/events/events.service';

@Controller('wishes')
export class WishesController {
  constructor(
    private readonly wishesService: WishesService,
    private readonly usersService: UsersService,
    private readonly eventsService: EventsService,
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
  @Delete(':id')
  async remove(@Request() req, @Param('id') id: number, @Body() createWishDto: CreateWishDto) {
    const user = await this.usersService.findByEmail(req.user.user.email);
    return await this.wishesService.remove(user, id, createWishDto);
  }
}
