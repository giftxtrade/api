import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Request, HttpStatus, HttpException } from '@nestjs/common';
import { ParticipantsService } from './participants.service';
import { CreateParticipantDto } from './dto/create-participant.dto';
import { UpdateParticipantDto } from './dto/update-participant.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { UsersService } from 'src/users/users.service';

@Controller('participants')
export class ParticipantsController {
  constructor(
    private readonly participantsService: ParticipantsService,
    private readonly usersService: UsersService,
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
}
