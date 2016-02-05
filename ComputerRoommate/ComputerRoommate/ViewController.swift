//
//  ViewController.swift
//  ComputerRoommate
//
//  Created by Dan Kurtz on 1/22/16.
//  Copyright © 2016 Dan Kurtz. All rights reserved.
//

import Cocoa
import Foundation

class ViewController: NSViewController {
    
    @IBAction func selectProfile(sender: NSButton) {
        let json = NSBundle.mainBundle().pathForResource(sender.title, ofType: "json")
        
        var resourcesParts = json!.characters.split("/").map(String.init)
        resourcesParts.removeAtIndex(resourcesParts.count-1)
        let resourcesPath = resourcesParts.joinWithSeparator("/")
        
        let path = NSBundle.mainBundle().pathForResource("roommate-cli", ofType: nil)
        
        let task = NSTask()
        task.launchPath = path
        task.arguments = ["-service="+json!, "-resources=/"+resourcesPath+"/"]
        
        let filepath = "/tmp/roommate.log"
        if (!NSFileManager.defaultManager().createFileAtPath(filepath, contents: nil, attributes: nil)) {
            print("could not create file")
        } else if let theHandle = NSFileHandle(forWritingAtPath: filepath) {
            theHandle.writeData("Starting roommate log...\n".dataUsingEncoding(NSUTF8StringEncoding)!)
            task.standardOutput = theHandle
            task.standardError = theHandle
            
            task.launch()
            //theHandle.closeFile()
        } else {
            print("could not open file")
        }
        
        NSApplication.sharedApplication().keyWindow?.close()
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.view.wantsLayer = true
        self.view.layer!.backgroundColor = NSColor.whiteColor().CGColor
        // Do any additional setup after loading the view.
    }

    override var representedObject: AnyObject? {
        didSet {
        // Update the view, if already loaded.
        }
    }


}

