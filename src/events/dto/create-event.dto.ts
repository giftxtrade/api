import { CreateParticipantDto } from "src/participants/dto/create-participant.dto";

export class CreateEventDto {
  name: string;
  description: string;
  budget: number;
  invitationMessage: string;
  drawAt: Date;
  closeAt: Date;
  participants: CreateParticipantDto[]
}
