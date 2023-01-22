import axios from 'axios'
import { 
    Command, 
    OptionValues 
} from 'commander'


interface ITranslation {
    translatedText: string
    sourceLanguage: string
}

interface ITranslate {
    translate(): Promise<ITranslation>
}

class Translate implements ITranslate {
    private _apiKey: string
    private _to: string
    private _text: string
    private _program: Command
    private _optValues: OptionValues

    constructor() {
        this._program = new Command()
        // create arguments here and pass them to the program
        this._optValues = (
            this._program
            .version('0.0.1')
            .requiredOption('-t, --to <to>', 'Translate to language')
            .requiredOption('-s, --text <text>', 'Source text')
            .option('-k, --api-key <key>', 'DeepL API Key. If you do not want to provide Api key here - use env variable DEEPL_API_KEY')
            .option('-p, --pro-api', 'Use free DeepL API (default: false)')
            .parse(process.argv)
            ).opts()
        // check if we have api key. We can't procceed without it
        (!this._optValues.apiKey) 
            ? this._apiKey = this.verifyApiKey() 
            : this._apiKey = this._optValues.apiKey
        // set all necessary properties which leter be used for translate func
        this._to = this._optValues.to
        this._text = this._optValues.text
    }

    private verifyApiKey() {
        /**
         * if client does not use --api-key arg, try to check for env variable and throw if not found
         */
        const apiKey = process.env.DEEPL_API_KEY
        
        if (!apiKey) {
            throw new Error('No Deepl Api found! Please set the DEEPL_API_KEY environment variable or use flag --api-key')
        }

        return apiKey
    }
    
    // if proApi arg is not set, assume that user wants to use free api
    private generateUrlScheme() {
        return this._optValues.proApi
            ? `https://api.deepl.com`
            : 'https://api-free.deepl.com'
    }

    public async translate(textToTranslate: string | null = null): Promise<ITranslation> {
        const payload = {
            auth_key: this._apiKey,
            text: textToTranslate ?? this._text,
            target_lang: this._to
        }
        
        const deepLApi = axios.create(
            {
                baseURL: this.generateUrlScheme()
            }
        )

        const result = await deepLApi.post(
            'v2/translate', payload
        )
        
        return {
            translatedText: result.data.translations[0].text,
            sourceLanguage: result.data.translations[0].detected_source_language
        }
    }
    // expose all options to the outside
    public get options() {
        return this._optValues
    }

    public set apiKey(key: string) {
        this._apiKey = key
    }
}

async function main() {
    try {
        const translate = new Translate()
        const translt = await translate.translate()
        console.table(translt)
    } catch (err) {
        console.error("[ERROR] ::", err.message)
    }
}

(async () => await main())()
