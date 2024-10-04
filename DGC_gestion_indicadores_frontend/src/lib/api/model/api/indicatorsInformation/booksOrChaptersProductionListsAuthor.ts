import type { Career } from "../career"

export interface PostBooksOrChaptersProductionListsAuthorCareersRequest {
    books_or_chapters_production_list_id: number
    author_id: number
    careers: number[]
}

export interface BooksOrChaptersProductionListsAuthorsCareersJoined {
    author_id: number
    author: string
    careers: Career[]
}

