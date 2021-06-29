import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpException, HttpStatus } from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { EventsService } from 'src/events/events.service';
import { UsersService } from 'src/users/users.service';
import { DrawsService } from './draws.service';

@Controller('draws')
export class DrawsController {
  constructor(
    private readonly drawsService: DrawsService,
    private readonly usersService: UsersService,
    private readonly eventsService: EventsService,
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

  @Get()
  findAll() {
    return this.drawsService.findAll();
  }
}
