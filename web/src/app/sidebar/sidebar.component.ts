import { Component, Input, OnInit } from '@angular/core';

import { site } from '../site';

@Component({
    selector: 'app-sidebar',
    templateUrl: 'sidebar.component.html',
    styleUrls: ['sidebar.component.scss'],
})
export class SidebarComponent implements OnInit {
    @Input() selectedSite: site;
    @Input() newSite: site;

    constructor() {}

    ngOnInit() {}
}
