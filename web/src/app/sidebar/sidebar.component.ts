import {
  Component,
  Input,
  OnInit
} from '@angular/core';

import { Site } from '../site'

@Component({
  selector: 'app-sidebar',
  templateUrl: 'sidebar.component.html',
  styleUrls: ['sidebar.component.scss'],
})

export class SidebarComponent implements OnInit {
  @Input() selectedSite: Site;
  @Input() newSite: Site;

  constructor() { }

  ngOnInit() { }

}
