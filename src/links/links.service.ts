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
    return await this.linksRepository.findOne({ code });
  }

  async findByEvent(event: Event): Promise<Link> {
    return await this.linksRepository
      .createQueryBuilder('l')
      .innerJoin('l.event', 'e')
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
    if (expirationDate)
      link.expirationDate = expirationDate;
    return await link.save();
  }
}
