export type Housing = House | Apartment;

interface Common {
    location: string;
    bedrooms: number;
    bathrooms: number;
    size: number;
    price: number;
    images: string[];
    parkingSpaces: number;
}

export interface House extends Common {
    garden: boolean;
    pool: boolean;
    yearBuilt?: number;
}

export interface Apartment {
    balcony: boolean;
    parkingSpaces: number;
    petFriendly: boolean;
}
