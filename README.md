# RDV Data Export to Volanta

This CLI utility converts the Royal Dutch Virtual Data exports to a format that Volanta can import.

## Limitations

- All flights treated as IFR
- All flights treated as Scheduled
- All flights treated as flown Offline
- Aircraft Registrations are all set to `PH-RDV`
- Block Fuel and Flight Fuel will be the same value
- Non-landing events are not imported

## Installing

Download the appropriate executable from the `Releases` section of this repository.

## Usage

### 1. Setup The Filesystem

Create a new working folder where you plan to perform this conversion. You need to create two folders: `./source` and `./output`.

Place the `reports.json` and `reservations.json` files into the `./source` directory.

You will also place the folder of flight data in the `./source` directory unmodified.

Your working directory should look something like this:

```text
/
├─ output/
├─ source/
│  ├─ 12345/
│  │  ├─ fipo.json
│  │  ├─ flightwinds.json
│  │  ├─ rawdata.json
│  │  ├─ rawdata.zip
│  ├─ 67890/
│  │  ├─ fipo.json
│  │  ├─ flightwinds.json
│  │  ├─ rawdata.json
│  │  ├─ rawdata.zip
│  ├─ reports.json
│  ├─ reservations.json
```

### 2. Generate Aircraft Mapping Template

Before you can run the converter, you will need to generate an aircraft mapping template that will map the detected aircraft strings from your simulator to the ICAO codes that Volanta expects.

To do this, from your working directory, run the executable program: `./rdv2volanta mappings`. This should generate a `mappings.json` in your working directory that looks something like this:

```json
{
  "Boeing 737-8K2NG  KLM 'New Livery' PH-BXZ Goldstar Textures": "",
  "PMDG MD-11 HD KLM PH-KCB": ""
}
```

You will need to manually update this file and add the ICAO code for each aircraft in the blank space. An example of a completed mapping file looks like this:

```json
{
  "Boeing 737-8K2NG  KLM 'New Livery' PH-BXZ Goldstar Textures": "B738",
  "PMDG MD-11 HD KLM PH-KCB": "MD11"
}
```

### 3. Run Converter

Now you run the convert itself by running: `./rdv2volanta convert`.

This may take a few minutes to run depending on how many flights need to be converted. The converted flight flights will be saved in the `./output` folder and can be directly imported into Volanta.

## Troubleshooting

### Error parsing aircraft

If you have very early flights from the first few months of RDV, there may be missing aircraft information on those flight reports. This is to be expected, but can be corrected manually after importing to Volanta.

### Error loading rawdata file

Make sure you placed the flight data in the correct directory structure as described in step 1.
