/* Ridgeline — personal seed data for Travis Johnson.
 *
 * Sources:
 *  - Strava monthly recap emails (mail@update.strava.com) — actual totals Apr–Sep 2024,
 *    plus Jun 2023–Mar 2024 estimated from the recap bar charts (scale verified
 *    against the April/June/July actuals).
 *  - Strava club emails: Cuyamaca Epic Strava Team Race (race day 2026-06-20),
 *    Lobster Ride, "Social SATURDAY" 8:15 AM group ride.
 *  - Top sports per Strava: cycling (dominant), yoga, running.
 */
var SEED = {
  profile: {
    name: "Travis Johnson",
    email: "travcjohnson@gmail.com",
    location: "San Diego, CA",
    stravaAthleteId: 9337896,
    primarySport: "Gravel / MTB cycling",
    secondarySports: ["Yoga", "Running"],
  },

  targetEvent: {
    name: "Cuyamaca Epic Team Race",
    date: "2026-06-20",
    kind: "Gravel/MTB team race",
    note: "Strava club race out of Cuyamaca. Social SATURDAY crew, 8:15 AM starts.",
  },

  // Monthly Strava totals. estimated:true rows are derived from recap chart bars.
  stravaMonthly: [
    { month: "2023-06", hours: 43, miles: 660, elevFt: 36000, estimated: true },
    { month: "2023-07", hours: 8,  miles: 120, elevFt: 6500,  estimated: true },
    { month: "2023-08", hours: 55, miles: 840, elevFt: 46000, estimated: true },
    { month: "2023-09", hours: 50, miles: 770, elevFt: 42000, estimated: true },
    { month: "2023-10", hours: 26, miles: 400, elevFt: 22000, estimated: true },
    { month: "2023-11", hours: 26, miles: 400, elevFt: 22000, estimated: true },
    { month: "2023-12", hours: 23, miles: 350, elevFt: 19000, estimated: true },
    { month: "2024-01", hours: 32, miles: 490, elevFt: 27000, estimated: true },
    { month: "2024-02", hours: 30, miles: 465, elevFt: 25500, estimated: true },
    { month: "2024-03", hours: 44, miles: 675, elevFt: 37000, estimated: true },
    { month: "2024-04", hours: 64, miles: 953, elevFt: 57772, activeDays: 22 },
    { month: "2024-05", hours: 73, miles: 1131, elevFt: 61499, activeDays: 25, pr: true },
    { month: "2024-06", hours: 32, miles: 549, elevFt: 29856 },
    { month: "2024-07", hours: 30, miles: 484, elevFt: 31742, activeDays: 12 },
    { month: "2024-08", hours: 34, miles: 520, elevFt: 31863 },
    { month: "2024-09", hours: 15, miles: 202, elevFt: 11345, activeDays: 7 },
  ],

  // Daily habits chosen for an endurance athlete + founder. Editable in-app.
  defaultHabits: [
    { id: "h-train",    name: "Ride or train",        emoji: "🚴", why: "Consistency beats heroics — Cuyamaca rewards base miles." },
    { id: "h-sleep",    name: "In bed by 10:30",       emoji: "🛏️", why: "Recovery is made in bed. Drives tomorrow's green light." },
    { id: "h-sunlight", name: "Morning sunlight 10m",  emoji: "🌅", why: "Anchors circadian rhythm; better sleep onset." },
    { id: "h-mobility", name: "Mobility / yoga 10m",   emoji: "🧘", why: "Your #2 Strava sport — keeps the engine supple." },
    { id: "h-alcohol",  name: "Alcohol-free day",      emoji: "🚫", why: "Single biggest lever on recovery score." },
    { id: "h-protein",  name: "Protein target hit",    emoji: "🥩", why: "Rebuild what the climbs tear down." },
  ],
};
