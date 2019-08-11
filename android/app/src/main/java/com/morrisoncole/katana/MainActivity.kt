package com.morrisoncole.katana

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import com.squareup.wire.GrpcClient
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.Job
import kotlinx.coroutines.launch
import okhttp3.OkHttpClient
import okhttp3.Protocol
import java.util.concurrent.TimeUnit

class MainActivity : AppCompatActivity() {

    private var networkJob = Job()
    private val coroutineScope = CoroutineScope(networkJob + Dispatchers.Main)

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        requestAsync()
    }

    private fun requestAsync() {
        coroutineScope.launch {
            Log.println(Log.ERROR, "blag", "Making request...")
            val definitionRequest = DefinitionRequest("かたな")

            val grpcClient = GrpcClient.Builder()
                .client(OkHttpClient.Builder()
                    .protocols(listOf(Protocol.H2_PRIOR_KNOWLEDGE))
                    .readTimeout(60, TimeUnit.MINUTES)
                    .writeTimeout(60, TimeUnit.MINUTES)
                    .callTimeout(60, TimeUnit.MINUTES)
                    .build())
                .baseUrl("http://192.168.1.219:50051")
                .build()
            val dictionaryService = grpcClient.create(Dictionary::class)
            val response = dictionaryService.RequestDefinition(definitionRequest)
            Log.println(Log.ERROR, "blag", response.definition[0])
        }
    }
}
