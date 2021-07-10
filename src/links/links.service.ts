import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import Link from './entity/link.entity';
import { generate } from 'randomstring';
import { Event } from 'src/events/entities/event.entity';

@Injectable()
export class LinksService {
  private readonly codeLen = 15;

  constructor(
    @InjectRepository(Link)
    private readonly linksRepository: Repository<Link>,
  ) { }

  async findOne(code: string): Promise<Link> {
    return await this.linksRepository
      .createQueryBuilder('l')
      .where('l.code = :code AND expirationDate > CURRENT_DATE()', {
        code: `${code}`
      })
      .getOne();
  }

  /**
   * Find link with a join on event
   */
  async findOneWithEvent(code: string): Promise<Link> {
    return await this.linksRepository
      .createQueryBuilder('l')
      .leftJoinAndSelect('l.event', 'e')
      .where('l.code = :code', { code: `${code}` })
      .getOne();
  }

  async findByEvent(event: Event): Promise<Link> {
    return await this.linksRepository
      .createQueryBuilder('l')
      .where('l.eventId = :eventId', { eventId: event.id })
      .getOne();
  }

  async generateValidCode(): Promise<string> {
    let code = generate(this.codeLen);
    let link = await this.findOne(code);
    while (link) {
      code = generate(this.codeLen);
      link = await this.findOne(code);
    }
    return code;
  }

  async create(event: Event, expirationDate?: Date): Promise<Link> {
    const link = new Link();
    link.code = await this.generateValidCode();
    link.event = event;
    if (expirationDate)
      link.expirationDate = expirationDate;
    return await link.save();
  }
}
