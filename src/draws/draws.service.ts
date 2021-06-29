import { Injectable } from '@nestjs/common';
import { CreateDrawDto } from './dto/create-draw.dto';
import { UpdateDrawDto } from './dto/update-draw.dto';

@Injectable()
export class DrawsService {
  create(createDrawDto: CreateDrawDto) {
    return 'This action adds a new draw';
  }

  findAll() {
    return `This action returns all draws`;
  }

  getDraws() {
    const participants = [];
    const N = participants.length;

    for (let i = 0; i < N; i++) {
      const drawn = new Set<number>();
      const pairs = new Map<number, number>();
      const start = i + 1;

      for (let j = 0; j < N; j++) {
        const index = j + 1;
        let draw = start + index;
        if (draw == index)
          draw++;
        if (draw > N)
          draw %= N;
        if (draw === 0)
          draw++;

        if (drawn.has(draw)) {
          pairs.clear();
          break;
        }
        drawn.add(draw);
        pairs.set(index, draw);
      }
      if (pairs.size == 0)
        continue;
      participants.push(pairs);
    }

    participants.forEach(p => {
      let str = "";
      let i = 0;
      p.forEach((v, k) => {
        str += `${k}->${v}${i == (p.size - 1) ? '' : ', '}`;
        i++;
      })
      console.log(str);
    });
  }
}
