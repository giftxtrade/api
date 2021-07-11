import { SENDGRID } from '../../auth-tokens.json'
import { Event } from 'src/events/entities/event.entity';
import { User } from 'src/users/entities/user.entity';
import { Participant } from 'src/participants/entities/participant.entity';

const sgMail = require('@sendgrid/mail')

export const mailingTemplates = {
  namesDrawn: 'd-957406e64b0a4a0286838b56fdb20e5e',
  newParticipant: 'd-0241b0c49d7f4757b0a1381118ab81c2'
}

export function sendMail(to: string, subject: string, templateId: string, templateData: any) {
  sgMail.setApiKey(SENDGRID.API_KEY);
  return sgMail
    .send({
      to: to,
      from: 'giftxtrade@giftxtrade.com',
      subject: subject,
      dynamic_template_data: templateData,
      template_id: templateId
    });
}

export function newParticipantMail(user: User, event: Event, newParticipant: Participant) {
  const subject = `${newParticipant.name} Has Joined ${event.name} - GiftTrade`;

  sendMail(user.email, subject, mailingTemplates.newParticipant, {
    year: new Date().getFullYear().toString(),
    user: user,
    event: event,
    participant: newParticipant
  })
    .then((res: any) => { })
    .catch((res: any) => { });
}