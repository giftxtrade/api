/* Do not change, this code is generated from Golang structs */


export interface Participant {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    created_by: User;
    modified_by: User;
    email: string;
    nickname?: string;
    address?: string;
    organizer: boolean;
    participates: boolean;
    accepted: boolean;
    event: Event;
    user?: User;
}
export interface User {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    email: string;
    name: string;
    imageUrl?: string;
    isActive: boolean;
}
export interface Event {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    created_by: User;
    modified_by: User;
    name: string;
    description?: string;
    budget: number;
    inviteMessage?: string;
    drawAt: Date;
    closeAt: Date;
    slug: string;
    participants?: Participant[];
}