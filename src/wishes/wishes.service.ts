import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { CreateWishDto } from './dto/create-wish.dto';
import { UpdateWishDto } from './dto/update-wish.dto';
import { EventsService } from 'src/events/events.service';
import { ProductsService } from 'src/products/products.service';
import { ParticipantsService } from 'src/participants/participants.service';
import { User } from 'src/users/entities/user.entity';
import { Wish } from 'src/wishes/entities/wish.entity';
import { Product } from 'src/products/entities/product.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Event } from 'src/events/entities/event.entity';

@Injectable()
export class WishesService {
  constructor(
    @InjectRepository(Wish)
    private readonly wishRepository: Repository<Wish>,
    private readonly eventsService: EventsService,
    private readonly productsService: ProductsService,

    private readonly participantsService: ParticipantsService,
  ) {}

  async create(user: User, createWishDto: CreateWishDto) {
    const { event, participant, product } = await this.getWishItems(
      user,
      createWishDto.eventId,
      createWishDto.productId,
      createWishDto.participantId,
    );

    const wish = new Wish();
    wish.user = user;
    wish.event = event;
    wish.participant = participant;
    wish.product = product;
    return await wish.save();
  }

  findAllByUserEvent(user: User, event: Event): Promise<Wish[]> {
    return this.wishRepository
      .createQueryBuilder('w')
      .leftJoinAndSelect('w.product', 'products')
      .where('w.userId = :userId AND w.eventId = :eventId', {
        userId: user.id,
        eventId: event.id,
      })
      .orderBy('w.id', 'DESC')
      .getMany();
  }

  async findOneByUserProductEvent(user: User, product: Product, event: Event) {
    return this.wishRepository
      .createQueryBuilder('w')
      .where(
        'w.userId = :userId AND w.eventId = :eventId AND w.productId = :productId',
        {
          userId: user.id,
          eventId: event.id,
          productId: product.id,
        },
      )
      .getOne();
  }

  async remove(user: User, createWishDto: CreateWishDto) {
    const { product, event, participant } = await this.getWishItems(
      user,
      createWishDto.eventId,
      createWishDto.productId,
      createWishDto.participantId,
    );

    if (!product || !event || !participant) {
      throw new HttpException(
        {
          message: 'Something went wrong',
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    const wish = await this.findOneByUserProductEvent(user, product, event);
    if (!wish) {
      throw new HttpException(
        {
          message: 'Could not delete wish item',
        },
        HttpStatus.BAD_REQUEST,
      );
    }
    return await wish.remove();
  }

  private async getWishItems(
    user: User,
    eventId: number,
    productId: number,
    participantId: number,
  ) {
    const event = await this.eventsService.findOneForUser(eventId, user);
    const product = await this.productsService.findOne(productId);

    if (!event || !product) {
      throw new HttpException(
        {
          message: 'Could not find event or product',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    const participant = await this.participantsService.findByEventAndUser(
      event,
      user,
    );
    if (!participant) {
      if (participant.organizer && !participant.participates) {
        throw new HttpException(
          {
            message:
              'You are not a participant for this group. Swith your status to participant if you want a wishlist',
          },
          HttpStatus.BAD_REQUEST,
        );
      }

      throw new HttpException(
        {
          message: 'You are not part of this event',
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    if (participant.id !== participantId) {
      throw new HttpException(
        {
          message: 'Invalid participant id',
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    return { event, participant, product };
  }
}
