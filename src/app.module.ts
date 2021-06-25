import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { UsersModule } from './users/users.module';
import { GoogleModule } from './google/google.module';
import { AuthModule } from './auth/auth.module';
import { ProductsModule } from './products/products.module';
import { CategoriesModule } from './categories/categories.module';
import { EventsModule } from './events/events.module';
import { ServeStaticModule } from '@nestjs/serve-static';
import { join } from 'path';
import { ParticipantsModule } from './participants/participants.module';
import { WishesModule } from './wishes/wishes.module';
import { LinksModule } from './links/links.module';
import { DrawsModule } from './draws/draws.module';

@Module({
  imports: [
    TypeOrmModule.forRoot(),
    ServeStaticModule.forRoot({
      rootPath: join(__dirname, '..', 'uploads'),
      serveRoot: '/static',
    }),
    UsersModule,
    GoogleModule,
    AuthModule,
    ProductsModule,
    CategoriesModule,
    EventsModule,
    ParticipantsModule,
    WishesModule,
    LinksModule,
    DrawsModule,
  ],
  controllers: [AppController],
  providers: [],
})
export class AppModule {}
