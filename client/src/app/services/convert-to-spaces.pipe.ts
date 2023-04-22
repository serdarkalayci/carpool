import {Pipe, PipeTransform} from "@angular/core";

@Pipe({
  name: 'FormatDate'
})
export class FormatDatePipe implements PipeTransform {

  transform(value: string | undefined): string {
    return value!.slice(8, 10) + '-' + value!.slice(5, 7) + '-' + value!.slice(0, 4);
  }
}
