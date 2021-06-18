export class CreateEventDto {
  name: string;
  description: string;
  budget: number;
  invitationMessage: string;
  drawAt: Date;
  closeAt: Date;
  participants: { name: string, email: string }[]
}
