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

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateParticipantDto: UpdateParticipantDto) {
    return this.participantsService.update(+id, updateParticipantDto);
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':participantId')
  async remove(@Request() req, @Param('participantId') participantId: number) {
    const user = await this.usersService.findByEmail(req.user.user.email);

    const participant = await this.participantsService.findOneWithUser(participantId);
    if (!participant && participant.user.id !== user.id) {
      throw new HttpException({
        message: 'Could not delete participant'
      }, HttpStatus.BAD_REQUEST);
    }
    return await this.participantsService.remove(participantId);
  }
}
