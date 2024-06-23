import { EventCard } from "@/components/EventCard";
import { Title } from "@/components/Title";
import { EventModel } from "@/models";

export default function Home () {
    const events: EventModel[] = [
        {
            id: "1",
            name: "Desenvolvimento de software",
            organization: "Cubos",
            date: "2024-12-06T04:00:00.000Z",
            location: "Recife",
        },
        {
            id: "1",
            name: "Desenvolvimento de software",
            organization: "Cubos",
            date: "2024-12-06T04:00:00.000Z",
            location: "Recife",
        },
        {
            id: "1",
            name: "Desenvolvimento de software",
            organization: "Cubos",
            date: "2024-12-06T04:00:00.000Z",
            location: "Recife",
        }
    ]

    return (
        <main>
            <Title>
                Eventos disponíveis
            </Title>
            <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
                {events.map((event) => (
                    <EventCard key={event.id} event={event} />
                ))}
            </div>
        </main>

    );
}
